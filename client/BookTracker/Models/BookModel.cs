using System;

namespace BookTracker.Models
{
    public class BookModel
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

        public override String ToString()
        {
            return title + " by " + author;
        }
    }
}
