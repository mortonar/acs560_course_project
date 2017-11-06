using BookTracker.HelperClasses;

namespace BookTracker.ViewModels
{
    public class ToReadListViewModel : ObservableObject, IPageViewModel
    {
        public string Name
        {
            get { return "To Read"; }
        }
    }
}
