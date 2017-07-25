# SearchableList

	Package `searchablelist` implements the `list` package doubly linked list
	and extends it with search methods

    type SearchableList
        New () *SearchableList
        ContainsElement (t *Element) bool
        Contains (t *Element) bool
        ContainsValue (v interface{}) bool
        FindFirst (v interface{}) *Element
        FindLast (v interface{}) *Element
        FindAll (v interface{}) []*Element

    type Element
        func (e *Element) Next() *Element
        func (e *Element) Prev() *Element

    type List
        func (l *List) Back() *Element
        func (l *List) Front() *Element
        func (l *List) Init() *List
        func (l *List) InsertAfter(v interface{}, mark *Element) *Element
        func (l *List) InsertBefore(v interface{}, mark *Element) *Element
        func (l *List) Len() int
        func (l *List) MoveAfter(e, mark *Element)
        func (l *List) MoveBefore(e, mark *Element)
        func (l *List) MoveToBack(e *Element)
        func (l *List) MoveToFront(e *Element)
        func (l *List) PushBack(v interface{}) *Element
        func (l *List) PushBackList(other *List)
        func (l *List) PushFront(v interface{}) *Element
        func (l *List) PushFrontList(other *List)
        func (l *List) Remove(e *Element) interface{}
