package main

import (
	"fmt"
)

// With thanks to: https://examples.yourdictionary.com/anagram-examples.html

func ExampleAnagram() {
	fmt.Printf("%#v\n", Anagram("race", "care"))
	// Output: true
}

func ExampleAnagram_force_fail() {
	fmt.Printf("%#v\n", Anagram("this test", "return false"))
	// Output: false
}

func ExampleAnagram_tar_rat() {
	fmt.Printf("%#v\n", Anagram("Tar", "Rat"))
	// Output: true
}

func ExampleAnagram_arc_car() {
	fmt.Printf("%#v\n", Anagram("Arc", "Car"))
	// Output: true
}

func ExampleAnagram_elbow_below() {
	fmt.Printf("%#v\n", Anagram("Elbow", "Below"))
	// Output: true
}

func ExampleAnagram_state_taste() {
	fmt.Printf("%#v\n", Anagram("State", "Taste"))
	// Output: true
}

func ExampleAnagram_cider_cried() {
	fmt.Printf("%#v\n", Anagram("Cider", "Cried"))
	// Output: true
}

func ExampleAnagram_dusty_study() {
	fmt.Printf("%#v\n", Anagram("Dusty", "Study"))
	// Output: true
}

func ExampleAnagram_night_thing() {
	fmt.Printf("%#v\n", Anagram("Night", "Thing"))
	// Output: true
}

func ExampleAnagram_inch_chin() {
	fmt.Printf("%#v\n", Anagram("Inch", "Chin"))
	// Output: true
}

func ExampleAnagram_brag_grab() {
	fmt.Printf("%#v\n", Anagram("Brag", "Grab"))
	// Output: true
}

func ExampleAnagram_cat_act() {
	fmt.Printf("%#v\n", Anagram("Cat", "Act"))
	// Output: true
}

func ExampleAnagram_bored_robed() {
	fmt.Printf("%#v\n", Anagram("Bored", "Robed"))
	// Output: true
}

func ExampleAnagram_save_vase() {
	fmt.Printf("%#v\n", Anagram("Save", "Vase"))
	// Output: true
}

func ExampleAnagram_angel_glean() {
	fmt.Printf("%#v\n", Anagram("Angel", "Glean"))
	// Output: true
}

func ExampleAnagram_stressed_desserts() {
	fmt.Printf("%#v\n", Anagram("Stressed", "Desserts"))
	// Output: true
}

func ExampleAnagram_dormitory_dirty_room() {
	fmt.Printf("%#v\n", Anagram("Dormitory", "Dirty room"))
	// Output: true
}

func ExampleAnagram_school_master_the_classroom() {
	fmt.Printf("%#v\n", Anagram("School master", "The classroom"))
	// Output: true
}

func ExampleAnagram_conversation_voices_rant_on() {
	fmt.Printf("%#v\n", Anagram("Conversation", "Voices rant on"))
	// Output: true
}

func ExampleAnagram_listen_silent() {
	fmt.Printf("%#v\n", Anagram("Listen", "Silent"))
	// Output: true
}

func ExampleAnagram_astronomer_moon_starer() {
	fmt.Printf("%#v\n", Anagram("Astronomer", "Moon starer"))
	// Output: true
}

func ExampleAnagram_the_eyes_they_see() {
	fmt.Printf("%#v\n", Anagram("The eyes", "They see"))
	// Output: true
}

func ExampleAnagram_a_gentleman_elegant_man() {
	fmt.Printf("%#v\n", Anagram("A gentleman", "Elegant man"))
	// Output: true
}

func ExampleAnagram_funeral_real_fun() {
	fmt.Printf("%#v\n", Anagram("Funeral", "Real fun"))
	// Output: true
}

func ExampleAnagram_the_morse_code_here_comes_dots() {
	fmt.Printf("%#v\n", Anagram("The Morse Code", "Here come dots"))
	// Output: true
}

func ExampleAnagram_eleven_plus_two_twelve_plus_one() {
	fmt.Printf("%#v\n", Anagram("Eleven plus two", "Twelve plus one"))
	// Output: true
}

func ExampleAnagram_slot_machines_cash_lost_in_me() {
	fmt.Printf("%#v\n", Anagram("Slot machines", "Cash lost in me"))
	// Output: true
}

func ExampleAnagram_fourth_of_july_joyful_fourth() {
	fmt.Printf("%#v\n", Anagram("Fourth of July", "Joyful Fourth"))
	// Output: true
}
