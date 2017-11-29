using BookTracker.HelperClasses;
using BookTracker.Messaging.Requests;
using BookTracker.Models;
using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Input;

namespace BookTracker.ViewModels
{
    public class BookViewModel : ObservableObject, ICommand, IPageViewModel
    {
        private BookModel bookModel;

        public BookViewModel(BookModel bookModel)
        {
            this.bookModel = bookModel;
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
            Debug.WriteLine("BookViewModel::Execute with " + parameter);
            String shelfName = (String)parameter;
            addBook(shelfName);
        }

        public void Update()
        {
            Debug.WriteLine("BookViewModel::Update");
        }

        private void addBook(String shelfName)
        {
            AddBook addBookReq = new AddBook
            {
                ShelfName = shelfName,
                Book = bookModel
            };
            Base req = new Base
            {
                Action = "AddBook",
                Payload = addBookReq
            };
            ServerProxy.Instance.sendRequest(req);
        }
    }
}
