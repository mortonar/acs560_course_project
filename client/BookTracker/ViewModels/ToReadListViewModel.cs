using BookTracker.HelperClasses;
using BookTracker.Models;

namespace BookTracker.ViewModels
{
    public class ToReadListViewModel : ObservableObject, IPageViewModel
    {
        public string Name
        {
            get { return "To Read"; }
        }

        public BookListModel BookListModel
        {
            get
            {
                BookListModel list = new BookListModel();
                list.AddBook(new BookModel("1984", "George Orwell", "B076FD7NKD"));
                list.AddBook(new BookModel("Scepters", "L.E. Modesitt, Jr.", "B06Y156FL4"));
                return list;
            }
        }
    }
}
