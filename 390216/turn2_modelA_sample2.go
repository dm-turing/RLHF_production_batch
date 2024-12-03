package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type SocialGraph struct {
	users map[int]map[int]struct{} // adj map
	mu    sync.Mutex
}

func NewSocialGraph() *SocialGraph {
	return &SocialGraph{
		users: make(map[int]map[int]struct{}),
	}
}

func (sg *SocialGraph) AddFriendship(user1, user2 int) {
	sg.mu.Lock()
	defer sg.mu.Unlock()

	if _, exists := sg.users[user1]; !exists {
		sg.users[user1] = make(map[int]struct{})
	}

	if _, exists := sg.users[user2]; !exists {
		sg.users[user2] = make(map[int]struct{})
	}

	sg.users[user1][user2] = struct{}{}
	sg.users[user2][user1] = struct{}{}
}

func predictLinks(sg *SocialGraph, numFriends int) {
	var wg sync.WaitGroup

	numUsers := len(sg.users)

	for i := 0; i < numUsers; i++ {
		for j := i + 1; j < numUsers; j++ {
			if _, isFriend := sg.users[i][j]; isFriend {
				continue
			}

			wg.Add(1)
			go func(user1, user2 int) {
				defer wg.Done()
				_, friend1 := sg.users[user1]
				_, friend2 := sg.users[user2]
				if !friend1 || !friend2 {
					return
				}

				numCommonFriends := len(intersection(sg.users[user1], sg.users[user2]))

				if numCommonFriends > numFriends {
					fmt.Println("Predicted Link:", user1, user2)
				}
			}(i, j)
		}
	}

	wg.Wait()
}

func intersection(m1, m2 map[int]struct{}) map[int]struct{} {
	var intersect map[int]struct{}
	intersect = make(map[int]struct{})
	for key, _ := range m1 {
		if _, found := m2[key]; found {
			intersect[key] = struct{}{}
		}
	}

	return intersect
}

func main() {
	sg := NewSocialGraph()

	// Simulate some friendships
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		for j := i + 1; j < 10; j++ {
			if rand.Intn(2) == 0 {
				sg.AddFriendship(i, j)
			}
		}
	}

	// Predict links based on common friends
	predictLinks(sg, 1)
}
