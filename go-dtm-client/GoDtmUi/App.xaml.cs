using Avalonia;
using Avalonia.Markup.Xaml;

namespace GoDtmUI
{
    public class App : Application
    {
        public override void Initialize()
        {
            AvaloniaXamlLoader.Load(this);
        }
   }
}