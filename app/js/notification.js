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
      var notifications = this.state.list.map(function(value){
        console.log(value.mail);
        return (
          <tr>
            <td>{ value.notificationType }</td>
            <td>{ value.mail.destination.join(", ") }</td>
            <td>{ value.mail.source }</td>
            <td>{ value.CreatedAt }</td>
            <td></td>
          </tr>
        )
      });
    return (<tbody>{ notifications }</tbody>);
  }
  else {
    return (<tbody><tr><td colSpan='4'></td></tr></tbody>)
    }
  },
});

ReactDOM.render(<NotificationComponent/>, document.getElementById('content'));
