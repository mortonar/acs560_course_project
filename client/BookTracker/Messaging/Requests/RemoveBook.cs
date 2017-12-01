using BookTracker.Models;

namespace BookTracker.Messaging.Requests
{
    public class RemoveBook
    {
        public BookModel Book { get; set; }
        public string ShelfName { get; set; }
    }
}
