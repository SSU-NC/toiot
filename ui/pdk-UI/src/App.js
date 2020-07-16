import React, { useImperativeHandle, Component } from 'react';
import { BrowserRouter as Router, Route } from 'react-router-dom';
import Nav from "./Navigation";
import RegisterSensor from "./components/RegisterSensor"
import RegisterNode from "./components/RegisterNode"


class App extends Component {
  constructor(props){
    super(props);
    this.state = {
      sensorlist: []
      // node: [],
      // rasp: []
    }
  }
  componentDidMount() {
    this.getSensorList();
  }

  getSensorList() {
    var url = 'http://220.70.2.160:8080/sensor/info';

    fetch(url)
      .then(res => res.json())
      .then(data => this.setState(
        {sensorlist: data}
      ))
      // .then(response => console.log('Success:', JSON.stringify(response)))
      .catch(error => console.error('Error:', error));
  }

  render() {
    return (
      <div>
        {/* <RegisterSensor></RegisterSensor> */}
        <RegisterNode sensorList={this.state.sensorlist}></RegisterNode>
      </div>
    );
  }
}

/*
function App() {
  return (
    <div>
        <Router>
          <div>
            <Nav></Nav>
            <div className='container pt-4 mt-4'> 
              <Route exact path="/" render={} />
              <Route path="/sensor" render={} />
              <Route path="/kibana" component={} />
            </div>
          </div>
        </Router>
      </div>
  );
}*/
/*
컴포넌트에 전달할 속성이 있을 경우 render 사용
  * <Route exact path="/" render={() => <CardContainer location='cards.json' member={true}/>} />
속성이 없다면 component 사용
  * <Route path="/about" component={About} />
*/
export default App;