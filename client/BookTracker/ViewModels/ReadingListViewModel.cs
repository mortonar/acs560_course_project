using BookTracker.HelperClasses;

namespace BookTracker.ViewModels
{
    public class ReadingListViewModel : ObservableObject, IPageViewModel
    {
        public string Name
        {
            get { return "Reading"; }
        }
    }
}
