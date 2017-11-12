using BookTracker.HelperClasses;
using BookTracker.Models;

namespace BookTracker.ViewModels
{
    public class ReadingListViewModel : ObservableObject, IPageViewModel
    {
        public string Name
        {
            get { return "Reading"; }
        }

        public BookListModel BookListModel
        {
            get
            {
                BookListModel list = new BookListModel();
                list.AddBook(new BookModel("Dracula", "Bram Stoker", "684176"));
                list.AddBook(new BookModel("Darkness", "L.E. Modesitt, Jr.", "B06Y156FL4"));
                return list;
            }
        }
    }
}
