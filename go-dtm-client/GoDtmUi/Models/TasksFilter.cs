namespace GoDtmUI.Models
{
    public class TasksFilter
    {
        private const string TYPE_DEFAULT = "default";
        private const string TYPE_ASIGNEE = "asignee";
        private const string TYPE_CREATOR = "creator";
        private const string TYPE_STATUS = "status";
        private const string TYPE_START = "start";
        private const string TYPE_END = "end";
        
        private Status _status;
        public Status Status
        {
            get => _status;
            set
            {
                _status = value;
                Type = TYPE_STATUS;
                Param = _status.Id;
            }
        }

        private User _asignee;
        public User Asignee
        {
            get => _asignee;
            set
            {
                _asignee = value;
                Type = TYPE_ASIGNEE;
                Param = _asignee.Id;
            }
        }

        private User _creator;
        public User Creator
        {
            get => _creator;
            set
            {
                _creator = value;
                Type = TYPE_CREATOR;
                Param = _creator.Id;
            }
        }

        private int _startDate;
        public int StartDate
        {
            get => _startDate;
            set
            {
                _startDate = value;
                Type = TYPE_START;
                Param = _startDate;
            }
        }

        private int _endDate;
        public int EndDate
        {
            get => _endDate;
            set
            {
                _endDate = value;
                Type = TYPE_END;
                Param = _endDate;
            }
        }
        
        public int Offset { get; set; }
        public int Limit { get; set; }
        
        public string Type { get; set; }
        public int Param { get; private set; }


        public TasksFilter()
        {
            Status = new Status();
            Asignee = new User();
            Creator = new User();

            Type = TYPE_DEFAULT;
            Param = 0;
        }
    }
}