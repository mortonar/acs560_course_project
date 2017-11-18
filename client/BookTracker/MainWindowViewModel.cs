using BookTracker.HelperClasses;
using BookTracker.ViewModels;
using System.Collections.Generic;
using System.Linq;
using System.Windows.Input;

namespace BookTracker
{
    public class MainWindowViewModel : ObservableObject
    {
        #region Fields

        private ICommand _changePageCommand;
        private IPageViewModel _currentPageViewModel;
        private List<IPageViewModel> _pageViewModels;
        #endregion

        public MainWindowViewModel()
        {
            // Add available pages
            PageViewModels.Add(new HomeViewModel());
            // TODO - Make this ViewModel either login or logout, depending on the current status of the session
            PageViewModels.Add(new LoginViewModel());
            PageViewModels.Add(new BookSearchViewModel());
            PageViewModels.Add(new ToReadListViewModel());
            PageViewModels.Add(new ReadingListViewModel());
            PageViewModels.Add(new ReadListViewModel());

            // Set starting page
            CurrentPageViewModel = PageViewModels[0];
        }

        #region Properties / Commands

        public ICommand ChangePageCommand
        {
            get
            {
                if (_changePageCommand == null)
                {
                    _changePageCommand = new RelayCommand(
                        p => ChangeViewModel((IPageViewModel)p),
                        p => p is IPageViewModel);
                }

                return _changePageCommand;
            }
        }

        public List<IPageViewModel> PageViewModels
        {
            get
            {
                if (_pageViewModels == null)
                    _pageViewModels = new List<IPageViewModel>();

                return _pageViewModels;
            }
        }

        public IPageViewModel CurrentPageViewModel
        {
            get
            {
                return _currentPageViewModel;
            }
            set
            {
                if (_currentPageViewModel != value)
                {
                    _currentPageViewModel = value;
                    OnPropertyChanged("CurrentPageViewModel");
                }
            }
        }
        #endregion

        #region Methods

        private void ChangeViewModel(IPageViewModel viewModel)
        {
            if (!PageViewModels.Contains(viewModel))
            {
                PageViewModels.Add(viewModel);
            }

            if (viewModel is ToReadListViewModel || viewModel is ReadingListViewModel || viewModel is ReadListViewModel)
            {
                viewModel.Update();
            }

            CurrentPageViewModel = PageViewModels
                .FirstOrDefault(vm => vm == viewModel);
        }
        #endregion
    }
}
