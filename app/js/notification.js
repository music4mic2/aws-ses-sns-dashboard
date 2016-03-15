var NotificationComponent = React.createClass({
  getInitialState: function() {
    return { list: {} };
  },
  componentDidMount: function() {
  var _this = this;

  //TODO
  $.get( "http://localhost:8000/notifications", { page: 1, limit: 2 }, function(data){})
    .done(function(data) {
      _this.setState( { list: data })
    });
  },
  render: function() {
    if ( this.state.list.length > 0 ) {
      var notifications = this.state.list.map(function(value, index){
        var date = new Date(value.CreatedAt);
        return (
          <tr key={index}>
            <td>{value.notificationType == "Delivery" ? <span className="label label-primary">{value.notificationType}</span> : <span className="label label-danger">{value.notificationType}</span>}<br /><small className="text-muted">{ date.toUTCString() }</small></td>
            <td>{ value.mail.destination.join(", ") }</td>
            <td>{ value.mail.source }</td>
            <td>{ value.bounce.bounceSubType || "" }</td>
          </tr>
        )
      });
    return (
        <table className="table table-striped">
          <thead>
          <tr>
            <th>Notificación</th>
            <th>Email</th>
            <th>Fuente</th>
            <th>Detalle</th>
          </tr>
          </thead>
          <tbody>
            { notifications }
          </tbody>
        </table>
        );
  }
  else {
    return (<span></span>)
    }
  },
});

ReactDOM.render(<NotificationComponent/>, document.getElementById('content'));
