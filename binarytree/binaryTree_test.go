package binarytree

import (
	"reflect"
	"testing"
)

func TestBinaryTree_Insert(t *testing.T) {
	type fields struct {
		Root *Node
	}
	type args struct {
		value int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   BinaryTree
	}{
		{
			name: "Should_AddToTheRoot_When_NilRoot",
			fields: fields{
				Root: nil,
			},
			args: args{
				value: 1,
			},
			want: BinaryTree{Root: &Node{Value: 1}},
		},
		{
			name: "Should_NotAddNewNode_When_SameValueAsRoot",
			fields: fields{
				Root: &Node{Value: 1},
			},
			args: args{
				value: 1,
			},
			want: BinaryTree{Root: &Node{Value: 1}},
		},
		{
			name: "Should_AddNewNodeOnTheRight_When_ValueGreaterThanRoot",
			fields: fields{
				Root: &Node{Value: 1},
			},
			args: args{
				value: 2,
			},
			want: BinaryTree{
				Root: &Node{
					Value: 1,
					Right: &Node{Value: 2},
				},
			},
		},
		{
			name: "Should_AddNewNodeOnTheRightNode_When_ValueGreaterThanRootRight",
			fields: fields{
				Root: &Node{Value: 1, Right: &Node{Value: 2}},
			},
			args: args{
				value: 3,
			},
			want: BinaryTree{
				Root: &Node{
					Value: 1,
					Right: &Node{Value: 2, Right: &Node{Value: 3}},
				},
			},
		},
		{
			name: "Should_AddNodeToTheLeft_When_ValueLesserThanRoot",
			fields: fields{
				Root: &Node{Value: 2},
			},
			args: args{
				value: 1,
			},
			want: BinaryTree{
				Root: &Node{
					Value: 2,
					Left:  &Node{Value: 1},
				},
			},
		},
		{
			name: "Should_AddNodeOnTheLeftNode_When_ValueLesserThanRootLeft",
			fields: fields{
				Root: &Node{Value: 3, Left: &Node{Value: 2}},
			},
			args: args{
				value: 1,
			},
			want: BinaryTree{
				Root: &Node{
					Value: 3,
					Left:  &Node{Value: 2, Left: &Node{Value: 1}},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bt := &BinaryTree{
				Root: tt.fields.Root,
			}
			bt.Insert(tt.args.value)
			if !reflect.DeepEqual(bt.Root, tt.want.Root) {
				t.Errorf("Insert() = %+v, want %+v", bt.Root, tt.want.Root)
			}
		})
	}
}

func TestBinaryTree_Contains(t *testing.T) {
	type fields struct {
		Root *Node
	}
	type args struct {
		value int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Should_ReturnFalse_When_Nil_Root",
			fields: fields{
				Root: nil,
			},
			args: args{
				value: 1,
			},
			want: false,
		},
		{
			name: "Should_ReturnTrue_When_ValueIsTheSameAsRoot",
			fields: fields{
				Root: &Node{Value: 1},
			},
			args: args{
				value: 1,
			},
			want: true,
		},
		{
			name: "Should_ReturnFalse_When_ValueIsLesserThanRootAndLeftIsNil",
			fields: fields{
				Root: &Node{Value: 2},
			},
			args: args{
				value: 1,
			},
			want: false,
		},
		{
			name: "Should_ReturnTrue_When_ValueIsTheSameAsRootLeft",
			fields: fields{
				Root: &Node{Value: 2, Left: &Node{Value: 1}},
			},
			args: args{
				value: 1,
			},
			want: true,
		},
		{
			name: "Should_ReturnFalse_When_ValueIsGreaterThanRootAndRightIsNil",
			fields: fields{
				Root: &Node{Value: 1},
			},
			args: args{
				value: 2,
			},
			want: false,
		},
		{
			name: "Should_ReturnTrue_When_ValueIsTheSameAsRootRight",
			fields: fields{
				Root: &Node{Value: 1, Right: &Node{Value: 2}},
			},
			args: args{
				value: 2,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bt := &BinaryTree{
				Root: tt.fields.Root,
			}
			if got := bt.Contains(tt.args.value); got != tt.want {
				t.Errorf("BinaryTree.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryTree_PrintLtr(t *testing.T) {
	type fields struct {
		Root *Node
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Should_PrintFromLeftToRight",
			fields: fields{
				Root: &Node{Value: 7, Left: &Node{Value: 3}, Right: &Node{Value: 10}},
			},
			want: "3 7 10",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bt := &BinaryTree{
				Root: tt.fields.Root,
			}
			if got := bt.PrintLtr(); got != tt.want {
				t.Errorf("BinaryTree.PrintLtr() = %v, want %v", got, tt.want)
			}
		})
	}
}
