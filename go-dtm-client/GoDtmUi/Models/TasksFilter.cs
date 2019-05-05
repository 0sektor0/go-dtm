namespace GoDtmUI.Models
{
    public class TasksFilter
    {
        public string AsigneeName { get; set; }
        public string CreatorName { get; set; }
        public string StatusName { get; set; }
        public int StartDate { get; set; }
        public int EndDate { get; set; }
        public int Offset { get; set; }
        public int Limit { get; set; }
    }
}