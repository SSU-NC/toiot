import React, { Component } from 'react';

// form : https://getbootstrap.com/docs/4.0/components/forms/?
// add, delete input : https://www.cluemediator.com/add-or-remove-input-fields-dynamically-with-reactjs
class RegisterSensor extends Component {

    // handle input change
    /*handleInputChange = (e, index) => {
        
    };*/
    
    // handle click event of the Remove button
    handleAddClick = index => {
       alert("+");
    };
    
    // handle click event of the Add button
    handleRemoveClick = () => {
        alert("-");
    };
    render() {
        return (
            <>
            <button type="button" class="btn btn-primary btn-lg" data-toggle="modal" data-target="#myModal">register sensor</button>
            <div class="modal fade" id="myModal" tabindex="-1" role="dialog" aria-labelledby="myModalLabel">
                <div class="modal-dialog" role="document">
                    <div class="modal-content">
                        <div class="modal-header"> 
                            <h4 class="modal-title" id="myModalLabel">Register sensor</h4>
                            <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">×</span></button>
                        </div>
                        <form>
                            <div class="modal-body">
                                <div class="form-group">
                                    <label for="sensor_name">Sensor name</label>
                                    <input type="text" class="form-control" id="sensor_name" placeholder="name"/>
                                    <div class="invalid-feedback">
                                        This sensor name is already exist.
                                    </div>
                                </div>
                                
                                <div class="form-group">
                                    <label for="">value name</label>
                                    <input type="text" class="form-control" id="value" placeholder="Enter value name" onChange={this.handleInputChange}/>
                                </div>
                                 <button type="button" class="btn btn-primary" onClick={this.handleAddClick}>+</button>

                                 <button type="button" class="btn btn-primary " onClick={this.handleRemoveClick}>-</button>
                               
                            </div>
                            <div class="modal-footer">
                                <button type="submit" class="btn btn-primary" data-dismiss="modal">Register</button>
                                <button type="button" class="btn btn-default" data-dismiss="modal">Cancel</button>
                            </div>
                        </form>
                    </div>
                </div>
            </div>
            </>
        );
    }
}


export default RegisterSensor;