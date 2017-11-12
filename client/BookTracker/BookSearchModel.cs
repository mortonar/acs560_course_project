using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace BookTracker
{
    public class BookSearchModel
    {
        private String _author;
        private String _title;

        public BookSearchModel()
        {
            _author = "George Orwell";
            _title = "1984";
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
