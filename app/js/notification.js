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
    var source = this.props.source;
    var type = this.props.type;
    this.props.fetchList(page, email, source, type);
  }
});

var FilterComponent = React.createClass({
  render: function() {
    return (
      <div className="row">
        <div className="col-sm-4 m-b-xs">
          <select name="type" id="type" className="input-sm form-control input-s-sm inline">
            <option value="">Todos</option>
            <option value="Delivery">Delivery</option>
            <option value="Bounce">Bounce</option>
          </select>
        </div>
        <div className="col-sm-3 m-b-xs">
          <input type="email" placeholder="Email" className="input-sm form-control" id="email"/>
        </div>
        <div className="col-sm-3 m-b-xs">
          <input type="source" placeholder="Fuente" className="input-sm form-control" id="source"/>
        </div>
        <div className="col-sm-2 pull-right">
         <button type="button" className="btn btn-sm btn-primary" onClick={this.clickHandler}>Buscar</button>
        </div>
      </div>
    )
  },
   clickHandler: function(){
     var email = $("#email").val();
     var source = $("#source").val();
     var type = $("#type").val();
     var page = this.props.page;
     this.props.fetchList(page, type, email, source);
   }
});

var NotificationComponent = React.createClass({
  getInitialState: function() {
    return { 
      list: {}, 
      page: 1, 
      type: null, 
      email: null,
      source: null,
    };
  },
  fetchList: function(page, type, email, source){
    var _this = this;
    var page = page || 1;
    var type =  type || "";
    var email = email || "";
    var source = source || "";

    var url = "http://localhost:8000/dashboard"

    $.ajax({
      method: "POST",
      url: url,
      withCredentials: true,
      data: { page: page, type: type, email: email, source: source },
      success: function( data ) {
        _this.setState({
          list: data,
          page: page,
          type: type,
          email: email,
          source: source
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
            <td>{value.notificationType == "Delivery" ? <span className="label label-primary">{value.notificationType}</span> : <span className="label label-danger">{value.notificationType}</span>}<br /><small className="text-muted">{ date.toLocaleString() }</small></td>
            <td>{ value.mail.destination.join(", ") }</td>
            <td><a href={'https://s3.amazonaws.com/beetrack-shared/emails/' +  value.mail.message_id} sorce target="_blank"><i class="fa fa-search" aria-hidden="true"></i></a></td>
            <td>{ value.mail.source }</td>
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
                <th>Notificación</th>
                <th>Email</th>
                <th>Fuente</th>
                <th>Contenido</th>
              </tr>
            </thead>
            <tbody>
              { notifications }
            </tbody>
          </table>
          <PaginationComponent fetchList={this.fetchList} page={this.state.page} email={this.state.email} type={this.state.type} source={this.state.source}/>
        </div>
      </div>
    );
  }
  else {
    return (
        <div>
          <FilterComponent fetchList={this.fetchList} page={1}/>
          <br />
          <div className="alert alert-info">No existen mas notificaciones.</div>
        </div>
      )
    }
  },
});

ReactDOM.render(<NotificationComponent/>, document.getElementById('content'));
