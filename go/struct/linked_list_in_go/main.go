// This is the practice of linked list in golang
// Ref: https://medium.com/backendarmy/linked-lists-in-go-f7a7b27a03b9

package main

import (
	"fmt"
)

type song struct {
	name   string
	artist string
	next   *song
}

type playlist struct {
	name       string
	head       *song
	nowPlaying *song
}

func createPlayList(name string) *playlist { // what is the type error?
	return &playlist{
		name: name,
	}
}

/*
1->2->x
   |
   cur
*/
func (p *playlist) addSong(name, artist string) error {
	s := &song{
		name:   name,
		artist: artist,
	}

	if p.head == nil {
		p.head = s
	} else {
		currentNode := p.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = s
	}
	return nil
}

/*
head
|
1 -> 2 -> 3 -> x
|
cur

*/
func (p *playlist) showAllSongs() error {
	currentNode := p.head
	// check if empty
	if currentNode == nil {
		fmt.Println("playlist is empty")
	}
	for currentNode != nil {
		// order of print?
		fmt.Printf("%+v\n", *currentNode)
		currentNode = currentNode.next
	}

	return nil
}

func (p *playlist) startPlaying() *song {
	p.nowPlaying = p.head
	return p.nowPlaying
}

func (p *playlist) nextSong() *song {
	if p.nowPlaying.next == nil {
		p.nowPlaying = p.head
		return p.nowPlaying
	}
	p.nowPlaying = p.nowPlaying.next
	return p.nowPlaying
}

func (p *playlist) showInfo() {
	fmt.Printf("Now playing: %s by %s\n", p.nowPlaying.name, p.nowPlaying.artist)
	fmt.Println()
}
func main() {
	// Create new list
	playlistName := "sleeping"
	myPlaylist := createPlayList(playlistName)
	fmt.Println("Creating playlist...")
	fmt.Println()

	// Add songs
	fmt.Println("Add songs to playlist...")
	myPlaylist.addSong("Ophelia", "The Lumineers") // addSong is a method
	myPlaylist.addSong("Shape of you", "Ed sheeran")
	myPlaylist.addSong("Stubborn Love", "The Lumineers")
	myPlaylist.addSong("Feels", "Calvin Harris")

	// show songs added
	myPlaylist.showAllSongs()
	fmt.Println()

	// Start playing
	myPlaylist.startPlaying()
	fmt.Printf("Now playing: %s by %s\n", myPlaylist.nowPlaying.name, myPlaylist.nowPlaying.artist)
	fmt.Println()

	// Next song
	myPlaylist.nextSong()
	fmt.Println("Changing next song...")
	myPlaylist.showInfo()
	// Next song
	myPlaylist.nextSong()
	fmt.Println("Changing next song...")
	myPlaylist.showInfo()

	// Next song
	myPlaylist.nextSong()
	fmt.Println("Changing next song...")
	myPlaylist.showInfo()

	// Next song
	myPlaylist.nextSong()
	fmt.Println("Changing next song...")
	myPlaylist.showInfo()
	//TODO: Implement function/method to remove songs
	//myPlaylist.removeSongs()
}
