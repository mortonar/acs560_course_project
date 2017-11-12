using BookTracker.HelperClasses;
using BookTracker.Models;
using System;

namespace BookTracker.ViewModels
{
    public class BookSearchResultsViewModel : ObservableObject, IPageViewModel
    {
        public string Name
        {
            get { return "BookSearchResults"; }
        }

        private BookListModel _searchResults;

        public event EventHandler CanExecuteChanged;

        public BookListModel BookListModel
        {
            get { return _searchResults; }
            set { _searchResults = value; }
        }

        public BookSearchResultsViewModel(BookListModel results)
        {
            _searchResults = results;
        }
    }
}
