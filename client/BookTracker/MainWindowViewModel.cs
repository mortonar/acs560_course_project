using BookTracker.HelperClasses;
using BookTracker.ViewModels;
using System;
using System.Collections.Generic;
using System.Diagnostics;
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
        // set of view models for which login isn't required to access
        private ISet<Type> loginExempt = new HashSet<Type>();
        #endregion

        public MainWindowViewModel()
        {
            // Add available pages
            // TODO - Make this ViewModel either login or logout, depending on the current status of the session
            PageViewModels.Add(new LoginViewModel());
            PageViewModels.Add(new ToReadListViewModel());
            PageViewModels.Add(new ReadingListViewModel());
            PageViewModels.Add(new ReadListViewModel());
            PageViewModels.Add(new BookSearchViewModel());

            // Set starting page
            CurrentPageViewModel = PageViewModels[0];

            // add any view models for which we aren't required login to view
            loginExempt.Add(typeof(LoginViewModel));
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
            if (!((App)App.Current).isLoggedIn() && !loginExempt.Contains(viewModel.GetType()))
            {
                // TODO display a messge in the ui to remind the user to log in
                Debug.WriteLine("Not logged in!");
                return;
            }
            if (!PageViewModels.Contains(viewModel))
            {
                PageViewModels.Add(viewModel);
            }            
            viewModel.Update();

            CurrentPageViewModel = PageViewModels
                .FirstOrDefault(vm => vm == viewModel);
        }
        #endregion
    }
}
