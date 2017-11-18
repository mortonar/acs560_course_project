using BookTracker.HelperClasses;
using BookTracker.Messaging.Requests;
using BookTracker.Models;
using Newtonsoft.Json;
using Newtonsoft.Json.Linq;
using System.Diagnostics;

namespace BookTracker.ViewModels
{
    public class ReadingListViewModel : ObservableObject, IPageViewModel
    {

        public ReadingListViewModel()
        {
            Debug.WriteLine("Getting the To Read Book List.");
            Update();
        }

        public void Update()
        {
            BookList bookList = new BookList
            {
                Name = Name
            };

            Base message = new Base
            {
                Action = "BookList",
                Payload = bookList
            };
            string response = ServerProxy.Instance.sendRequest(message);
            Debug.WriteLine("RESPONSE: " + response + "\n");
            Messaging.Responses.Base responseMsg = JsonConvert.DeserializeObject<Messaging.Responses.Base>(response);
            Debug.WriteLine("PAYLOD: " + responseMsg.Payload);
            Messaging.Responses.BookSearch searchResponse = (responseMsg.Payload as JObject).ToObject<Messaging.Responses.BookSearch>();
            _bookListModel = new BookListModel(searchResponse.Books);
        }

        public string Name
        {
            get { return "Reading"; }
        }

        private BookListModel _bookListModel;
        public BookListModel BookListModel
        {
            get { return _bookListModel; }
            set { _bookListModel = value; }
        }
    }
}
