var PaginationComponent = React.createClass({
  getDefaultProps: function() {
    page: 1
  },
  render: function() {
    return (
      <ul className="pagination">
        <li><a href= "javascript:void(0);" onClick={this.clickHandler.bind(this,this.props.page - 1)}>&laquo;</a></li>
        <li><a href= "javascript:void(0);" onClick={this.clickHandler.bind(this,this.props.page + 1)}>&raquo;</a></li>
      </ul>
    )
  },
  clickHandler: function(page, e) {
    $(".pagination").removeClass("active");
    $("li.pagination:eq("+ (page - 1) +")").addClass("active");
    var email = this.props.email;
    this.props.fetchList(page, email);
  }
});

var FilterComponent = React.createClass({
  render: function() {
    return (
      <div className="row">
        <div className="form-group">
          <div className="col-sm-4">
            <div className="input-group"><input type="email" placeholder="Email" className="input-sm form-control" id="email" required />
              <span className="input-group-btn">
                <button type="button" className="btn btn-sm btn-primary" onClick={this.clickHandler}>Buscar</button>
              </span>
            </div>
          </div>
        </div>
      </div>
    )
  },
   clickHandler: function(){
     var email = $("#email").val();
     var page = this.props.page;
     this.props.fetchList(page, email);
   }
});

var NotificationComponent = React.createClass({
  getInitialState: function() {
    return { 
      list: {}, 
      page: 1, 
      email: null
    };
  },
  fetchList: function(page, email){
    var _this = this;
    var page = page || 1;
    var email = email || "";

    var url = "http://localhost:8000/dashboard"

    $.ajax({
      method: "GET",
      url: url,
      withCredentials: true,
      data: { page: page, email: email },
      success: function( data ) {
        _this.setState({
          list: data,
          page: page,
          email: email
        });
      },
      beforeSend: function (xhr) {
        xhr.setRequestHeader("Authorization", "Basic " + btoa("admin:admin"));
      }
    });
  },
  componentDidMount: function() {
    this.fetchList();
  },
  render: function() {
    var list_length = this.state.list.length;
    if ( list_length > 0 ) {
      var notifications = this.state.list.map(function(value, index){
        var date = new Date(value.CreatedAt);
        return (
          <tr key={index}>
            <td>{value.MailID}</td>
            <td>{value.notificationType == "Delivery" ? <span className="label label-primary">{value.notificationType}</span> : <span className="label label-danger">{value.notificationType}</span>}<br /><small className="text-muted">{ date.toUTCString() }</small></td>
            <td>{ value.mail.destination.join(", ") }</td>
            <td>{ value.mail.source }</td>
            <td>{ value.bounce.bounceSubType || "" }</td>
          </tr>
        )
      });
    return (
      <div>
        <FilterComponent fetchList={this.fetchList} page={this.state.page}/>
        <div className="table-responsive">
          <table className="table table-striped">
            <thead>
              <tr>
                <th>id</th>
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
          <PaginationComponent fetchList={this.fetchList} page={this.state.page} email={this.state.email}/>
        </div>
      </div>
    );
  }
  else {
    return (<div className="alert alert-info">No existen mas notificaciones.</div>)
    }
  },
});
ReactDOM.render(<NotificationComponent/>, document.getElementById('content'));
