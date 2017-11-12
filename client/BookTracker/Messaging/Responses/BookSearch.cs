using BookTracker.Models;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace BookTracker.Messaging.Responses
{
    public class BookSearch
    {

        public List<BookModel> Books { get; set; }

    }
}
