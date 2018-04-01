var client = new Keen({
  projectId: '5368fa5436bf5a5623000000',
  readKey: '3f324dcb5636316d6865ab0ebbbbc725224c7f8f3e8899c7733439965d6d4a2c7f13bf7765458790bd50ec76b4361687f51cf626314585dc246bb51aeb455c0a1dd6ce77a993d9c953c5fc554d1d3530ca5d17bdc6d1333ef3d8146a990c79435bb2c7d936f259a22647a75407921056'
});

Keen.ready(function(){

  // Pageviews by browser

  var pageviews_timeline = new Keen.Dataviz()
    .el('#chart-01')
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
                pageviews_timeline.message("Could not request time monthly data");
                console.error(data);
                return;
            }
            console.info("data", data);
            pageviews_timeline
                .data(data)
                .sortGroups('desc')
                .render();
        },
        dataType: "json",
        error: function(e) {
            console.error(e);
            pageviews_timeline.message("Could not request time monthly data");
        }
    });


  // Pageviews by browser (pie)
  var pageviews_pie = new Keen.Dataviz()
    .el('#chart-02')
    .type('pie')
    .height(280)
    .title('Expenses')
    .prepare();

    $.ajax({
      type: "POST",
      url: "/api/v1/expenses",
      data: JSON.stringify({Directive: "list"}),
      success: function(data) {
          if (data.hasOwnProperty("reason") || data.hasOwnProperty("code")) {
              pageviews_pie.message("Could not request expenses data");
              console.error(data);
              return;
          }
          console.info("data", data);
          pageviews_pie
            .data(data)
            .sortGroups('desc')
            .render();
      },
      dataType: "json",
      error: function(e) {
          console.error(e);
          pageviews_pie.message("Could not request expenses data");
      }
    });


  // Impressions timeline

  var impressions_timeline = new Keen.Dataviz()
    .el('#chart-03')
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
                pageviews_pie.message("Could not request expenses data");
                console.error(data);
                return;
            }
            console.info("data", data);
            impressions_timeline
                .data(data)
                .sortGroups('desc')
                .render();
        },
        dataType: "json",
        error: function(e) {
            console.error(e);
            pageviews_pie.message("Could not request expenses data");
        }
    });

  // Impressions by device

  var impressions_by_device = new Keen.Dataviz()
    .el('#chart-04')
    .type('area')
    .height(280)
    .stacked(false)
    .title('Profit')
    .prepare();

  $.ajax({
    type: "POST",
    url: "/api/v1/profit",
    data: JSON.stringify({Directive: "list"}),
    success: function(data) {
        if (data.hasOwnProperty("reason") || data.hasOwnProperty("code")) {
            pageviews_pie.message("Could not request expenses data");
            console.error(data);
            return;
        }
        console.info("data", data);
        impressions_by_device
            .data(data)
            .sortGroups('desc')
            .render();
    },
    dataType: "json",
    error: function(e) {
        console.error(e);
        pageviews_pie.message("Could not request expenses data");
    }
  });
});
