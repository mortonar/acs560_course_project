using BookTracker.HelperClasses;
using System.Windows.Input;
using System;
using System.Diagnostics;
using BookTracker.Messaging.Requests;
using Newtonsoft.Json;
using BookTracker.Models;
using Newtonsoft.Json.Linq;

namespace BookTracker.ViewModels
{
    public class BookSearchViewModel : ObservableObject, ICommand, IPageViewModel
    {

        private BookSearchModel _bookSearchModel;
        public BookSearchModel BookSearchModel
        {
            get { return _bookSearchModel; }
            set { _bookSearchModel = value; }
        }

        public string Name
        {
            get { return "Search Book"; }
        }

        public event EventHandler CanExecuteChanged;

        public BookSearchViewModel()
        {
            _bookSearchModel = new BookSearchModel();
        }

        public bool CanExecute(object parameter)
        {
            return true;
        }

        public void Execute(object parameter)
        {
            BookSearch();
        }

        public void BookSearch()
        {
            Debug.Write("Searching for book with author | title: " +
                BookSearchModel.Author + " | " + BookSearchModel.Title);
            BookSearch bookSearch = new BookSearch
            {
                Author = _bookSearchModel.Author,
                Title = _bookSearchModel.Title
            };
            Base message = new Base
            {
                Action = "BookSearch",
                Payload = bookSearch
            };
            string response = ServerProxy.Instance.sendRequest(message);
            Debug.WriteLine("RESPONSE: " + response + "\n");
            Messaging.Responses.Base responseMsg = JsonConvert.DeserializeObject<Messaging.Responses.Base>(response);
            Debug.WriteLine("PAYLOD: " + responseMsg.Payload);
            Messaging.Responses.BookSearch searchResponse = (responseMsg.Payload as JObject).ToObject<Messaging.Responses.BookSearch>();
            ((App)App.Current).changeViewModel(new BookSearchResultsViewModel(new BookListModel(searchResponse.Books)));
        }

        public void Update() { }
    }
}
