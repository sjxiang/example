package main

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	
	"github.com/sjxiang/example/pb"
)

func convertShelf(shelf *Shelf) *pb.Shelf {
	return &pb.Shelf{
		Id:        shelf.ID,
		Theme:     shelf.Theme,
		Size:      shelf.Size,
		CreatedAt: timestamppb.New(shelf.CreateAt),
		UpdatedAt: timestamppb.New(shelf.UpdateAt),
	}
}

func convertShelves(shelves []*Shelf) []*pb.Shelf {
	result := make([]*pb.Shelf, 0, len(shelves))
	for _, shelf := range shelves {
		result = append(result, convertShelf(shelf))
	}

	return result
}

func convertBook(book *Book) *pb.Book {
	return &pb.Book{
		Id:        book.ID,
		Author:    book.Author,
		Title:     book.Title,
		CreatedAt: timestamppb.New(book.CreateAt),
		UpdatedAt: timestamppb.New(book.UpdateAt),
	}
}

func convertBooks(books []*Book) []*pb.Book {
	result := make([]*pb.Book, 0, len(books))
	for _, book := range books {
		result = append(result, convertBook(book))
	}

	return result
}

func convertBooksWithSize(books []*Book, size int) []*pb.Book {
	result := make([]*pb.Book, 0, size)
	
	for i := 0; i < size; i ++ {
		result = append(result, convertBook(books[i]))	
	}
	
	return result
}