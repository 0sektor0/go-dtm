using Newtonsoft.Json;



namespace GoDtmUI.Models
{
    public partial class Attachment
    {
        [JsonProperty("path")]
        public string Path { get; set; }

        [JsonProperty("creationDate")]
        public int CreationDate { get; set; }
    }
}