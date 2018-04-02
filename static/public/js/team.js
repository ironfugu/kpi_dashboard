Keen.ready(function(){

    var start = moment().subtract(1, "years");
    var end = moment();

    $('input[name="daterange"]').daterangepicker(
        {
            locale: {
                format: 'YYYY-MM-DD'
            },
            startDate: start,
            endDate: end,
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
            showGraphs(start, end);
        });
  // contribution by role
  var contribution_by_role = new Keen.Dataviz()
    .el('#contribution')
    .type('bar')
    .height(280)
    .stacked(true)
    .title('Contribution by role')
    .prepare();

  // Commits timeline
  var commits_timeline = new Keen.Dataviz()
    .el('#commits')
    .type('bar')
    .height(280)
    .stacked(true)
    .title('Commits')
    .prepare();

    function showGraphs(start, end){
        $.ajax({
          type: "POST",
          url: "/api/v1/contribution",
          data: JSON.stringify({Directive: "list", Params: [start, end]}),
          success: function(data) {
              if (data.hasOwnProperty("reason") || data.hasOwnProperty("code")) {
                  contribution_by_role.message("Could not request time monthly data");
                  console.error(data);
                  return;
              }
              console.info("data", data);
              contribution_by_role
                  .data(data)
                  .sortGroups('desc')
                  .render();
          },
          dataType: "json",
          error: function(e) {
              console.error(e);
              contribution_by_role.message("Could not request time monthly data");
          }
        });
        $.ajax({
            type: "POST",
            url: "/api/v1/commits",
            data: JSON.stringify({Directive: "list", Params: [start, end]}),
            success: function(data) {
                if (data.hasOwnProperty("reason") || data.hasOwnProperty("code")) {
                    commits_timeline.message("Could not request expenses data");
                    console.error(data);
                    return;
                }
                console.info("data", data);
                commits_timeline
                    .data(data)
                    .sortGroups('desc')
                    .render();
            },
            dataType: "json",
            error: function(e) {
                console.error(e);
                commits_timeline.message("Could not request expenses data");
            }
        });
    }
    showGraphs(start, end);
});
