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
		name:       name,
		head:       nil,
		nowPlaying: nil,
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
		p.nowPlaying = s
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
		fmt.Printf("%+v\n", *currentNode)
		currentNode = currentNode.next
	}

	fmt.Println()

	return nil
}

func (p *playlist) startPlaying() {
	p.nowPlaying = p.head
}

func (p *playlist) nextSong() {
	fmt.Println("Changing next song...")
	if p.nowPlaying.next == nil {
		p.nowPlaying = p.head
	}
	p.nowPlaying = p.nowPlaying.next
}
func (p *playlist) removeSong(target string) { // by linus in c : passed in node what if I want to passed in
	// deal with empty and nil playlist
	if p == nil || p.head == nil {
		fmt.Println("Empty or Nil playlist")
		return
	}

        fmt.Printf("Removing song: %q...\n", target)
	inHead := &p.head
	for (**inHead).name != target {
		inHead = &(*inHead).next
	}
	*inHead = (*inHead).next
	// declare pointer to pointer point to &p.head
	// traverse list
	// if target == current node
	// set next pointer to target.next
	return
}
func (p *playlist) showInfo() {
	if p == nil {
		fmt.Println("Nil playlist")
		return
	}
	if p.head == nil {
		fmt.Println("Empty playlist")
		return
	}
	fmt.Printf("Now playing: %s by %s\n", p.nowPlaying.name, p.nowPlaying.artist)
	fmt.Println()
}
func main() {
	// Create new list
	playlistName := "sleeping"
	myPlaylist := createPlayList(playlistName)
	fmt.Println("Creating playlist...")
	fmt.Println()

	myPlaylist.showInfo()
	// Add songs
	fmt.Println("Add songs to playlist...")
	myPlaylist.addSong("Ophelia", "The Lumineers") // addSong is a method
	myPlaylist.addSong("Shape of you", "Ed sheeran")
	myPlaylist.addSong("Stubborn Love", "The Lumineers")
	myPlaylist.addSong("Feels", "Calvin Harris")

	// show songs added
	err := myPlaylist.showAllSongs()
	if err != nil {
		fmt.Println("Error in showAllSongs...")
	}
	fmt.Println()

	// Start playing
	myPlaylist.startPlaying()
	myPlaylist.showInfo()

	// Next song
	myPlaylist.nextSong()
	myPlaylist.showInfo()
	// Next song
	myPlaylist.nextSong()
	myPlaylist.showInfo()

	// Next song
	myPlaylist.nextSong()
	myPlaylist.showInfo()

	// Next song
	myPlaylist.nextSong()
	myPlaylist.showInfo()
	//TODO: Try to handle errors of every operation

        // Remove song "Feels"
	myPlaylist.removeSong("Feels")
	err = myPlaylist.showAllSongs()
	if err != nil {
		fmt.Println("Error in showAllSongs...")
	}
        // Remove song "Shape of you"	
	myPlaylist.removeSong("Shape of you")

	err = myPlaylist.showAllSongs()
	if err != nil {
		fmt.Println("Error in showAllSongs...")
	}
        // Remove song "Stubborn Love"	
	myPlaylist.removeSong("Stubborn Love")

	err = myPlaylist.showAllSongs()
	if err != nil {
		fmt.Println("Error in showAllSongs...")
	}
        // Remove song "Ophelia"	
	myPlaylist.removeSong("Ophelia")

	err = myPlaylist.showAllSongs()
	if err != nil {
		fmt.Println("Error in showAllSongs...")
	}

        // Remove song "Ophelia" again from empty playlist, test robusity
        myPlaylist.removeSong("Ophelia")
        if err != nil {
            fmt.Println("Error in showAllSongs...")
        }
}
