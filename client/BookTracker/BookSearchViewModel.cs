using BookTracker.HelperClasses;

namespace BookTracker
{
    public class BookSearchViewModel : ObservableObject, IPageViewModel
    {
        public string Name
        {
            get { return "Search Book"; }
        }
    }
}
