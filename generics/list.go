package generics


import "github.com/google/go-cmp/cmp"

type List[T any] struct{
	Items []T
}

func (list *List[T]) Add(item T){
	list.Items = append(list.Items,item)
}

func (list *List[T]) InsertAt(index int,item T){
	// Our goals is move number between items of lists ===> 1 2 3 4 5 ---> 1 2 30 3 4 5 

	list.Items = append(list.Items, item) //  1 2 3 4 5 30
	copy(list.Items[index+1:],list.Items[index:]) // 1 2 3 3 4 5 
	list.Items[index] = item // 1 2 30 3 4 5
}

func (list *List[T]) RemoveAt(index int){
	// 1 2 3 4 5 --> 1 2 4 5
	list.Items = append(list.Items[:index], list.Items[index+1:]...)
	
}

func (list *List[T]) Remove(item T){
	// 1 2 3 4 5 --> 1 2 4 5
	index := list.Find(item)
	if index != -1 {
		list.RemoveAt(index)
	}
}

func (list *List[T]) Get(index int) T{
	return list.Items[index]
}


func (list *List[T]) Find(item T)int{
	for i, v := range list.Items{
		if cmp.Equal(v,item){
			return i
		}
	}
	return -1
}