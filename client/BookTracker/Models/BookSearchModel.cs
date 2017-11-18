using System;

namespace BookTracker.Models
{
    public class BookSearchModel
    {
        private String _author;
        private String _title;

        public BookSearchModel()
        {
            _author = "1984";
            _title = "George Orwell";
        }

        public String Author
        {
            get { return _author; }
            set { _author = value; }
        }

        public String Title
        {
            get { return _title; }
            set { _title = value; }
        }
    }

}
