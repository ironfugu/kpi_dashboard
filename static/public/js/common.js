window.Common = {
    Start: moment().subtract(1, "years"),
    End: moment(),
    daterange_change: function (callback) {
        $('input[name="daterange"]').daterangepicker(
            {
                locale: {
                    format: 'YYYY-MM-DD'
                },
                startDate: this.Start,
                endDate: this.End,
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
                this.Start = start;
                this.End = end;
                callback(start, end, label);
            });
    },
    projectsPrepare: function (onchange) {
        this.$projects = $("#projects");
        $.ajax({
            type: "POST",
            url: "/api/v1/projects",
            data: JSON.stringify({Directive: "list"}),
            success: function(data) {
                if (data.hasOwnProperty("reason") || data.hasOwnProperty("code")) {
                    console.error(data);
                    return;
                }
                console.info("data", data);
                data.forEach(function (value) {
                    this.$projects.append("<option value='"+value.id+"'>"+value.title+"</option>");
                }.bind(this));
                this.$projects.selectpicker();
                this.$projects.change(function () {
                    var projId = this.$projects.val();
                    onchange(projId);
                }.bind(this));
            }.bind(this),
            dataType: "json",
            error: function(e) {
                console.error(e);
            }
        });
    }
};