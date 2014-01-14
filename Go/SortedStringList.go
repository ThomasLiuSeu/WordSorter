package main

type List struct {
    length int
    middle *Node
    head *Node
}

type Node struct {
    next *Node
    value string
}

func (self *List) Add(value string) {
    if value == "" {
        return
    } else if self.head == nil {
        self.head = &Node{nil, ""}
        self.middle = self.head
    }

    if self.middle.value < value {
        self.AddAfterNode(self.middle, value)
    } else {
        self.AddAfterNode(self.head, value)
    }
}

func (self *List) AddAfterNode(node *Node, value string) {
    current := node
    middleIndex := self.length / 2
    newNode := &Node{nil, value}

    for i := 0; ; i++ {
        if current.next == nil {
            current.next = newNode
            break;
        } else if current.next.value > value {
            newNode.next = current.next
            current.next = newNode
            break
        }

        if (i == middleIndex) {
            self.middle = current
        }

        current = current.next
    }

    self.length++
}

func (self *List) ToSlice() (result []string) {
    result = make([]string, self.length)

    if self.length == 0 {
        return
    }

    i := 0
    for current := self.head.next; current != nil; current = current.next {
        result[i] = current.value
        i++
    }

    return
}
