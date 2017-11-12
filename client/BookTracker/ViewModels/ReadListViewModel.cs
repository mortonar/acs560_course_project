using BookTracker.HelperClasses;
using BookTracker.Models;

namespace BookTracker.ViewModels
{
    public class ReadListViewModel : ObservableObject, IPageViewModel
    {
        public string Name
        {
            get { return "Read"; }
        }

        public BookListModel BookListModel
        {
            get
            {
                BookListModel list = new BookListModel();
                list.AddBook(new BookModel("1984", "George Orwell", "B076FD7NKD"));
                list.AddBook(new BookModel("Dracula", "Bram Stoker", "684176"));
                list.AddBook(new BookModel("Legacies", "L.E. Modesitt, Jr.", "B06Y156FL4"));
                return list;
            }
        }
    }
}
