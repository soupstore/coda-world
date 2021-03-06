package model

import "github.com/google/uuid"

// ContainerID is a type-aliased string, often set to a uuid.
type ContainerID string

// Container is a common interface for any type of container that can store items in it.
type Container interface {
	PutItem(item *Item)
	RemoveItem(itemID ItemID)
	ID() ContainerID
	Items() map[ItemID]*Item
}

// BaseContainer is provided to be embedded in more complex containers.
type BaseContainer struct {
	id    ContainerID
	items map[ItemID]*Item
}

func (c *BaseContainer) ID() ContainerID {
	return c.id
}

func (c *BaseContainer) Items() map[ItemID]*Item {
	return c.items
}

// RoomContainer represents the floor of rooms items are dropped on to
type RoomContainer struct {
	BaseContainer
}

func NewRoomContainer() Container {
	return &RoomContainer{
		BaseContainer: BaseContainer{
			id:    ContainerID(uuid.New().String()),
			items: make(map[ItemID]*Item),
		},
	}
}

func (c *RoomContainer) PutItem(item *Item) {
	c.items[item.ID] = item
}

func (c *RoomContainer) RemoveItem(itemID ItemID) {
	delete(c.items, itemID)
}

// ItemContainer is the kind of container used in items like chests, etc...
type ItemContainer struct {
	BaseContainer
}

func NewItemContainer() Container {
	return &ItemContainer{
		BaseContainer: BaseContainer{
			id:    ContainerID(uuid.New().String()),
			items: make(map[ItemID]*Item),
		},
	}
}

func (c *ItemContainer) PutItem(item *Item) {
	// check for capacity and stuff
	c.items[item.ID] = item
}

func (c *ItemContainer) RemoveItem(itemID ItemID) {
	delete(c.items, itemID)
}

// CharacterContainer is a player's inventory.
type CharacterContainer struct {
	BaseContainer
}

func NewCharacterContainer() Container {
	return &CharacterContainer{
		BaseContainer: BaseContainer{
			id:    ContainerID(uuid.New().String()),
			items: make(map[ItemID]*Item),
		},
	}
}

func (c *CharacterContainer) PutItem(item *Item) {
	// check for capacity and stuff
	c.items[item.ID] = item
}

func (c *CharacterContainer) RemoveItem(itemID ItemID) {
	delete(c.items, itemID)
}
