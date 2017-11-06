using BookTracker.HelperClasses;

namespace BookTracker
{
    public class HomeViewModel : ObservableObject, IPageViewModel
    {
        public string Name
        {
            get { return "Home"; }
        }
    }
}
