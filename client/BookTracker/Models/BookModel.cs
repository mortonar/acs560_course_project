using BookTracker.ViewModels;
using System;
using System.Diagnostics;
using System.Windows.Input;
using System.Windows.Media;
using System.Windows.Media.Imaging;

namespace BookTracker.Models
{
    public class BookModel : ICommand
    {
        private String title;
        private String author;
        private String isbn13;
        private String description;
        private String thumbnailURL;
        private String imageURL;

        public BookModel(string title, string author, string isbn13)
        {
            this.title = title;
            this.author = author;
            this.isbn13 = isbn13;
        }

        public String Title
        {
            get { return title; }
            set { title = value; }
        }

        public String Author
        {
            get { return author; }
            set { author = value; }
        }

        public String Isbn13
        {
            get { return isbn13; }
            set { isbn13 = value; }
        }

        public String Description
        {
            get { return description;  }
            set { description = value; }
        }

        public String ThumbnailURL
        {
            get { return thumbnailURL; }
            set { thumbnailURL = value; }
        }

        public String ImageURL
        {
            get { return imageURL; }
            set { imageURL = value; }
        }

        public event EventHandler CanExecuteChanged;

        public bool CanExecute(object parameter)
        {
            return true;
        }

        public void Execute(object parameter)
        {
            Debug.WriteLine("Book model execute(" + this.ToString() + ")");
            if (parameter.Equals("Book"))
            {
                ((App)App.Current).changeViewModel(new BookViewModel(this));
            }
            else
            {
                ((App)App.Current).changeViewModel(new BookViewListsModel(this, (string)parameter));
            }
        }

        public String Listing
        {
            get
            {
                return ToString();
            }
        }

        public ImageSource Image
        {
            get { return getDisplayImage(); }
        }

        public override String ToString()
        {
            return title + " by " + author;
        }

        private ImageSource getDisplayImage()
        {
            
            String image = String.IsNullOrEmpty(imageURL) ? thumbnailURL : imageURL;
            if (String.IsNullOrEmpty(image))
            {
                image = "/BookTracker;component/Resources/CoverNotFound.bmp";

            }
            return new BitmapImage(new Uri(image, UriKind.RelativeOrAbsolute));
        }
    }
}
