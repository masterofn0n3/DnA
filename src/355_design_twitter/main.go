package main

import "container/heap"

type Twitter struct {
	userMap  map[int]map[int]struct{}
	tweetMap map[int]*Tweet
	time     int
}

type MaxHeap []*Tweet

func (m MaxHeap) Len() int            { return len(m) }
func (m MaxHeap) Less(i, j int) bool  { return m[i].time > m[j].time }
func (m MaxHeap) Swap(i, j int)       { m[i], m[j] = m[j], m[i] }
func (m *MaxHeap) Push(x interface{}) { *m = append(*m, x.(*Tweet)) }
func (m *MaxHeap) Pop() interface{} {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[:n-1]

	return x
}

type Tweet struct {
	time    int
	tweetId int
	next    *Tweet
	userId  int
}

func Constructor() Twitter {
	return Twitter{
		userMap:  map[int]map[int]struct{}{},
		tweetMap: map[int]*Tweet{},
		time:     0,
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	newTweet := &Tweet{
		tweetId: tweetId,
		userId:  userId,
		time:    this.time,
	}

	newTweet.next = this.tweetMap[userId]
	this.tweetMap[userId] = newTweet
	this.time += 1
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	followees := this.userMap[userId]

	h := MaxHeap([]*Tweet{})

	if this.tweetMap[userId] != nil {
		heap.Push(&h, this.tweetMap[userId])
	}

	for k := range followees {
		if this.tweetMap[k] != nil {
			heap.Push(&h, this.tweetMap[k])
		}
	}

	result := []int{}
	for len(result) < 10 && len(h) > 0 {
		pop := heap.Pop(&h).(*Tweet)
		result = append(result, pop.tweetId)
		if pop.next != nil {
			heap.Push(&h, pop.next)
		}
	}
	return result
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if this.userMap[followerId] == nil {
		this.userMap[followerId] = map[int]struct{}{}
	}
	this.userMap[followerId][followeeId] = struct{}{}
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	delete(this.userMap[followerId], followeeId)
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
