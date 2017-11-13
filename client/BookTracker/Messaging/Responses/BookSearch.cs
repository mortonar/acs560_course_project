using BookTracker.Models;
using System.Collections.Generic;

namespace BookTracker.Messaging.Responses
{
    public class BookSearch
    {

        public List<BookModel> Books { get; set; }

    }
}
