Keen.ready(function(){
    window.Common.daterange_change(showGraphs);
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
  window.Common.projectsPrepare(function (projId) {
        showGraphs(window.Common.Start, window.Common.End, projId);
  });
  function showGraphs(start, end, projID){
        $.ajax({
          type: "POST",
          url: "/api/v1/contribution",
          data: JSON.stringify({Directive: "list", Params: [start, end, projID]}),
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
            data: JSON.stringify({Directive: "list", Params: [start, end, projID]}),
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
    showGraphs(window.Common.Start, window.Common.End);
});
