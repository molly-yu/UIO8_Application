import React, {Component} from 'react';
import {Link} from 'react-router-dom';
import styled from 'styled-components';

const Styles = styled.div`
     .navbar{
        text-align:center;
        background: #2F80ED
        position: fixed;
        top: 0;
        width: 100%;
        z-index:10;
        margin:0;
        border:0;
    }
    .navbar-brand, .navbar-nav, .nav-link .navbar-collapse{
        color: #f2f2f2;
        font-size: 20 px;
        &:hover{
            color: #A4CCEB;
            border-bottom: 3px;
        }
    }
    .navbar a {
        float: left;
        display: block;
        color: #f2f2f2;
        text-align: center;
        padding: 14px 16px;
        text-decoration: none;
        font-size: 17px;
        border-bottom: 3px solid transparent;
        &:hover {
            border-bottom: 3px;
            color: #A4CCEB;
          }
          
          &.active {
            border-bottom: 3px;
            color: #74AAD5;
          }
      }   
      
`;


export default class Navbar extends Component{

    render(){
        return(
            <Styles>
            <nav className="navbar navbar-expand-lg">
                {/* <Link to="/" className="navbar-brand">ReBootUI</Link> */}
                <div className="collapse navbar-collapse">
                    {/* <ul className="navbar-nav mr-auto">
                     <li className="navbar-item">
                            <Link to="/" className="nav-link">Setup</Link>
                        </li>
                        <li className="navbar-item">
                            <Link to="/results" className="nav-link">Results</Link>
                        </li> 
                        
                    </ul> */}
                <a href="#setup">Setup</a>
                <a href="#results">Results</a>
                </div>
            </nav>
            </Styles>
        );
    }

}