Keen.ready(function(){

  // contribution by role

  var contribution_by_role = new Keen.Dataviz()
    .el('#contribution')
    .type('bar')
    .height(280)
    .stacked(true)
    .title('Contribution by role')
    .prepare();

  $.ajax({
        type: "POST",
        url: "/api/v1/time-monthly",
        data: JSON.stringify({Directive: "list"}),
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


    // Commits timeline
    var commits_timeline = new Keen.Dataviz()
        .el('#commits')
        .type('bar')
        .height(280)
        .stacked(true)
        .title('Commits')
        .prepare();


    $.ajax({
        type: "POST",
        url: "/api/v1/commits",
        data: JSON.stringify({Directive: "list"}),
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


});
