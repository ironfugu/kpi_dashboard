window.Common = {
    DEFAULT_START: moment().subtract(1, "years"),
    DEFAULT_END: moment(),
    daterange_change: function (callback) {
        $('input[name="daterange"]').daterangepicker(
            {
                locale: {
                    format: 'YYYY-MM-DD'
                },
                startDate: this.DEFAULT_START,
                endDate: this.DEFAULT_END,
                ranges: {
                    'Today': [moment(), moment()],
                    'Yesterday': [moment().subtract(1, 'days'), moment().subtract(1, 'days')],
                    'Last 7 Days': [moment().subtract(6, 'days'), moment()],
                    'Last 30 Days': [moment().subtract(29, 'days'), moment()],
                    'This Month': [moment().startOf('month'), moment().endOf('month')],
                    'Last Month': [moment().subtract(1, 'month').startOf('month'), moment().subtract(1, 'month').endOf('month')],
                    'Last 3 Months': [moment().subtract(3, 'month'), moment()],
                    'Last 6 Months': [moment().subtract(6, 'month'), moment()],
                    'Last 12 Months': [moment().subtract(12, 'month'), moment()]
                }
            },
            function(start, end, label) {
                callback(start, end, label);
            });
    }
};