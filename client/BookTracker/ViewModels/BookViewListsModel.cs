using BookTracker.HelperClasses;
using BookTracker.Messaging.Requests;
using BookTracker.Models;
using System;
using System.Diagnostics;
using System.Windows.Input;

namespace BookTracker.ViewModels
{
    public class BookViewListsModel : ObservableObject, ICommand, IPageViewModel
    {
        private BookModel bookModel;
        private string ShelfName;

        public BookViewListsModel(BookModel bookModel, string ShelfName)
        {
            this.bookModel = bookModel;

            // The shelf this book is currently on
            this.ShelfName = ShelfName;
        }

        public BookModel BookModel
        {
            get { return bookModel; }
            set { bookModel = value; }
        }

        public string Name
        {
            get { return "Book"; }

        }

        public event EventHandler CanExecuteChanged;

        public bool CanExecute(object parameter)
        {
            return true;
        }

        public void Execute(object parameter)
        {
            Debug.WriteLine("BookViewListsModel::Execute with " + parameter);
            String shelfName = (String)parameter;
            UpdateList(shelfName);
        }

        public void Update()
        {
            Debug.WriteLine("BookViewModel::Update");
        }

        private void UpdateList(String shelfName)
        {
            Base req;
            if (shelfName.Equals("RemoveBook"))
            {
                RemoveBook removeBookReq = new RemoveBook
                {
                    ShelfName = this.ShelfName,
                    Book = bookModel
                };
                req = new Base
                {
                    Action = "RemoveBook",
                    Payload = removeBookReq
                };
            }
            else
            {
                AddBook addBookReq = new AddBook
                {
                    ShelfName = shelfName,
                    Book = bookModel
                };
                req = new Base
                {
                    Action = "AddBook",
                    Payload = addBookReq
                };
            }
            ServerProxy.Instance.sendRequest(req);
        }
    }
}
