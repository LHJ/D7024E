package model

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
)

// IDLength is the static number of bytes in a KademliaID
const IDLength = 20

// KademliaID is the ID of files and node, represented by a fixed array of byte
type KademliaID [IDLength]byte

// NewKademliaID returns a new instance of a KademliaID based on the string input
func NewKademliaID(data []byte) *KademliaID {
	hash := sha1.Sum(data)
	slice := hash[0:20]

	newKademliaID := KademliaID{}
	for i := 0; i < IDLength; i++ {
		newKademliaID[i] = slice[i]
	}

	return &newKademliaID
}

//KademliaIDFromString convert an string ID into a KademliaID
func KademliaIDFromString(data string) *KademliaID {
	decoded, err := hex.DecodeString(data)
	if err != nil {
		log.Fatal("ERR" + err.Error())
	}

	newKademliaID := KademliaID{}
	for i := 0; i < IDLength; i++ {
		newKademliaID[i] = decoded[i]
	}

	return &newKademliaID
}

// KademliaIDFromBytes safely cast an variable array of bytes into a KademliaID
func KademliaIDFromBytes(unparsedID []byte) (id *KademliaID, err error) {
	if len(unparsedID) != IDLength {
		return nil, fmt.Errorf("invalid sized ID : '%s' of size %d>%d", unparsedID, len(unparsedID), IDLength)
	}
	var res = &KademliaID{}
	copy(res[:], unparsedID[:])
	return res, nil
}

// NewRandomKademliaID returns a new instance of a random KademliaID,
// change this to a better version if you like
func NewRandomKademliaID() *KademliaID {
	newKademliaID := KademliaID{}
	for i := 0; i < IDLength; i++ {
		newKademliaID[i] = uint8(rand.Intn(256))
	}
	return &newKademliaID
}

// Less returns true if kademliaID < otherKademliaID (bitwise)
func (kademliaID KademliaID) less(otherKademliaID *KademliaID) bool {
	result := false
	for i := 0; i < IDLength; i++ {
		if kademliaID[i] != otherKademliaID[i] {
			result = kademliaID[i] < otherKademliaID[i]
			break
		}
	}
	return result
}

// Equals returns true if kademliaID == otherKademliaID (bitwise)
func (kademliaID KademliaID) equals(otherKademliaID *KademliaID) bool {
	for i := 0; i < IDLength; i++ {
		if kademliaID[i] != otherKademliaID[i] {
			return false
		}
	}
	return true
}

// CalcDistance returns a new instance of a KademliaID that is built
// through a bitwise XOR operation betweeen kademliaID and target
func (kademliaID KademliaID) calcDistance(target *KademliaID) *KademliaID {
	result := KademliaID{}
	for i := 0; i < IDLength; i++ {
		result[i] = kademliaID[i] ^ target[i]
	}
	return &result
}

// String returns a simple string representation of a KademliaID
func (kademliaID *KademliaID) String() string {
	return hex.EncodeToString(kademliaID[0:IDLength])
}
