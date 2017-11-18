using System;
using System.Diagnostics;
using System.Windows.Input;

namespace BookTracker.Models
{
    public class BookModel : ICommand
    {
        private String title;
        private String author;
        private String isbn13;

        public BookModel(string title, string author, string isbn13)
        {
            this.title = title;
            this.author = author;
            this.isbn13 = isbn13;
        }

        public String Title
        {
            get { return title; }
            set { title = value; }
        }

        public String Author
        {
            get { return author; }
            set { author = value; }
        }

        public String Isbn13
        {
            get { return isbn13; }
            set { isbn13 = value; }
        }

        public event EventHandler CanExecuteChanged;

        public bool CanExecute(object parameter)
        {
            return true;
        }

        public void Execute(object parameter)
        {
            Debug.WriteLine("Book model execute(" + this.ToString() + ")");
        }

        public String Listing
        {
            get
            {
                return ToString();
            }
        }

        public override String ToString()
        {
            return title + " by " + author;
        }
    }
}
