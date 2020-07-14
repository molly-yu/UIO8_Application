import React, {Component} from 'react';
import {Link} from 'react-router-dom';
import axios from 'axios';
import PropTypes from 'prop-types';
import {connect } from 'react-redux';
import styled from 'styled-components';
import { Form, FormControl, InputGroup, Row, Col, Button, Table } from 'react-bootstrap';
import { createCamera } from '../actions/cameraActions';
const Styles = styled.div`
margin: 2em;
height:100vh;
width:100%;
 border:0;
 z-index: 5;
`;



class Cameras extends Component{
    constructor(props){
        super(props);
        this.state = {
            ip:'',
            user:'',
            pass:'',
        }
        this.onChange=this.onChange.bind(this);
        this.onSubmit= this.onSubmit.bind(this);
    }


    onChange(e){
        this.setState({[e.target.name]: e.target.value}); // set the state of the particular component
    }

    onSubmit(e){
        e.preventDefault();

        const camera = {
            ip: this.state.ip,
            user: this.state.user,
            pass: this.state.pass
        };
        this.props.createCamera(camera); // replaces fetch with createPost action
        
    }


    render(){
        return(
            <Styles>
            <div className="Cameras" id="cameras">
                <h2>Search: </h2>
                <Form onSubmit={this.onSubmit}>
                    <Form.Row>
                        <Col>
                        <Form.Label>Camera IP </Form.Label>
                        <InputGroup className="mb-3">
                            <InputGroup.Prepend >
                            <InputGroup.Text id="basic-addon3">
                                192.168.0.
                            </InputGroup.Text>
                            </InputGroup.Prepend>
                            <FormControl name="ip" onChange={this.onChange} value={this.state.ip} aria-describedby="basic-addon3" />
                        </InputGroup>
                        </Col>
                        <Col>
                        <Form.Group >
                            <Form.Label>Username</Form.Label>
                            <Form.Control name="user" onChange={this.onChange} value={this.state.user} placeholder="Username" />
                        </Form.Group>
                        </Col>
                        <Col>
                        <Form.Group >
                            <Form.Label>Password</Form.Label>
                            <Form.Control name="pass" onChange={this.onChange} value={this.state.pass} placeholder="Password" />
                        </Form.Group>
                        </Col>
                    </Form.Row>
                    

                    
                    <Button variant="primary" type="submit">
                        Add
                    </Button>
                    </Form>

            </div>
            </Styles>
        )
    }
}

Cameras.propTypes = {
    createCamera: PropTypes.func.isRequired
};

export default connect(null, {createCamera})(Cameras);