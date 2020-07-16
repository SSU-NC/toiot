import React, { Component } from 'react';
import Select from 'react-select';
// react-select : https://github.com/JedWatson/react-select

class RegisterNode extends Component {
    constructor(props) {
        super(props);
        this.state = {
            node_name: "",
            location: "",
            sensors:[]
        };

        this.handleNameChange = this.handleNameChange.bind(this);
        this.handleLocationChange = this.handleLocationChange.bind(this);
        this.handleSensorsChange = this.handleSensorsChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
    }

    handleNameChange(e) {
        this.setState({
            [e.target.name]: e.target.value
        });
    }
    handleLocationChange(e) {
        this.setState({
            [e.target.name]: e.target.value
        });
    }
    handleSensorsChange = sensors => {
        this.setState(
            { sensors }
        );
    }
    handleSubmit(e) {
        e.preventDefault();
      
        var url = 'http://220.70.2.160:8080/node/regist';
        var data = this.state;
        var sensor_uuid = data.sensors.map((val) => {
            return {uuid : val.uuid};
        });

        console.log(
            JSON.stringify({
            name: data.node_name,
            location: data.location,
            sensors: sensor_uuid
        }));

        fetch(url, {
        method: 'POST', // or 'PUT'
        body: JSON.stringify({
            name: data.node_name,
            location: data.location,
            sensors: sensor_uuid
        }),
        headers:{
            'Content-Type': 'application/json'
        }
        }).then(res => res.json())
        .then(response => console.log('Success:', JSON.stringify(response)))
        .catch(error => console.error('Error:', error));
    }

    isNodeNameValid = () =>{
        
    }

    render() {
        let sensorOptions = this.props.sensorList.map((val) =>{
            return {label : val.name, value: val.name, uuid: val.uuid}
        });

        return (
            <>
            <button type="button" class="btn btn-primary btn-lg" data-toggle="modal" data-target="#myModal">register node</button>
            <div class="modal fade" id="myModal" tabindex="-1" role="dialog" aria-labelledby="myModalLabel">
                <div class="modal-dialog" role="document">
                    <div class="modal-content">
                        <div class="modal-header">
                            <h4 class="modal-title" id="myModalLabel">Register node</h4>
                            <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">×</span></button>
                        </div>
                        <div class="modal-body">
                            <form >
                                <div class="form-group">
                                    <label for="node_name">Node name</label>
                                    <input type="text" class="form-control" name="node_name" placeholder="name" value={this.state.node_name} onChange={this.handleNameChange} />
                                    {/*<div class="invalid-feedback">
                                        This node name is already exist.
                                    </div>*/}
                                </div>
                                <div class="form-group">
                                    <label for="location">Location</label>
                                    <input type="text" class="form-control" name="location" placeholder="location" value={this.state.location} onChange={this.handleLocationChange}/>
                                </div>
                                <div class="form-group">
                                    <label for="select_sensor">Select sensors</label>
                                    <Select
                                        isMulti
                                        class="form-control"
                                        name="sensors"
                                        options={sensorOptions}
                                        className="basic-multi-select"
                                        classNamePrefix="select"
                                        value={this.state.sensors}
                                        onChange={this.handleSensorsChange}
                                    />
                                </div>
                                <div class="modal-footer">
                                    <button type="submit" class="btn btn-primary" data-dismiss="modal" onClick={this.handleSubmit}>Submit</button>
                                    <button type="reset" class="btn btn-default" data-dismiss="modal">Cancel</button>
                                </div>
                            </form>
                        </div>
                    </div>
                </div>
            </div>
            </>
        );
    }
}

export default RegisterNode;