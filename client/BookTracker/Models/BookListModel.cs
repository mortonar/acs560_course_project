using System.Collections;
using System.Collections.Generic;

namespace BookTracker.Models
{
    public class BookListModel
    {
        private List<BookModel> bookList;

        public BookListModel()
        {
            this.bookList = new List<BookModel>();
        }

        public BookListModel(List<BookModel> bookList)
        {
            this.bookList = bookList;
        }

        public List<BookModel> Booklist
        {
            get { return bookList; }
            set { bookList = value; }
        }

        public void AddBook(BookModel book)
        {
            bookList.Add(book);
        }

        public void RemoveBook(BookModel book)
        {
            bookList.Remove(book);
        }
    }
}
