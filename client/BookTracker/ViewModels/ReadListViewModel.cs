using BookTracker.HelperClasses;

namespace BookTracker.ViewModels
{
    public class ReadListViewModel : ObservableObject, IPageViewModel
    {
        public string Name
        {
            get { return "Read"; }
        }
    }
}
