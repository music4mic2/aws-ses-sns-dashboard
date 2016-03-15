var NotificationComponent = React.createClass({
  getInitialState: function() {
    return { list: {} };
  },
  componentDidMount: function() {
  var _this = this;
  //TODO
  $.get( "http://localhost:8000/notifications", function(data){})
    .done(function(data) {
      _this.setState( { list: data })
    });
  },
  render: function() {
    if ( this.state.list.length > 0 ) {
      var notifications = this.state.list.map(function(value, index){
        return (
          <tr key={index}>
            <td>{value.notificationType == "Delivery" ? <span className="label label-primary">{value.notificationType}</span> : <span className="label label-danger">{value.notificationType}</span>}</td>
            <td>{ value.mail.destination.join(", ") }</td>
            <td>{ value.mail.source }</td>
            <td>{ value.CreatedAt }</td>
            <td></td>
          </tr>
        )
      });
    return (
        <table className="table table-striped">
          <thead>
          <tr>
            <th>Notificación</th>
            <th>Email</th>
            <th>Descripción</th>
            <th>Fecha</th>
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
