using BookTracker.Models;

namespace BookTracker.Messaging.Requests
{
    public class AddBook
    {
        public BookModel Book { get; set; }
        public string ShelfName { get; set; }
    }
}
