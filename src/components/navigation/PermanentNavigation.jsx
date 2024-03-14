import * as React from 'react';
import Box from '@mui/material/Box';
import Drawer from '@mui/material/Drawer';
import CssBaseline from '@mui/material/CssBaseline';
import AppBar from '@mui/material/AppBar';
import Toolbar from '@mui/material/Toolbar';
import List from '@mui/material/List';
import Typography from '@mui/material/Typography';
import Divider from '@mui/material/Divider';
import ListItem from '@mui/material/ListItem';
import ListItemButton from '@mui/material/ListItemButton';
import ListItemIcon from '@mui/material/ListItemIcon';
import ListItemText from '@mui/material/ListItemText';
import InboxIcon from '@mui/icons-material/MoveToInbox';
import MailIcon from '@mui/icons-material/Mail';
import Field from '../field/OwnFiled.jsx';
import Price from '../price/price.jsx'
import './permanentNavigation.css'
import CollapseReg from '../collapse/collapse.tsx'
import SetButtonNav from './ButtonSetAcc.jsx';


const drawerWidth = 34;

export default function PermanentNavigation() {
  return (
    // <Box sx={{ display: 'flex' }}>
    //   {/* <CssBaseline /> */}
    //   <AppBar
    //     position="fixed"
    //     sx={{ width: `calc(100% - ${drawerWidth}px)`, ml: `${drawerWidth}px` }}
    //   >
    //     <Toolbar>
    //       <Typography variant="h6" noWrap component="div">
    //         Permanent drawer
    //       </Typography>
    //     </Toolbar>
    //   </AppBar>
    // <div className="wrapper">
    <>
      <Drawer
        sx={{
          width: '34%',
          'background-color': '#191016',
          flexShrink: 0,
          '& .MuiDrawer-paper': {
            width: '100%',
            boxSizing: 'border-box',
            position: 'relative',
            left: '0',
          'background-color': '#191016',
          'color': 'white',
          },
        }}
        variant="permanent"
        anchor="left"
      >
        <SetButtonNav/>
        <div className="logo">
        <svg width="264" height="75" viewBox="0 0 264 75" fill="none" xmlns="http://www.w3.org/2000/svg">
          <path d="M167.12 20.765C171.007 20.765 174.362 21.7367 177.185 23.68C180.045 25.5867 182.227 28.2267 183.73 31.6H176.03C173.793 28.6667 170.823 27.2 167.12 27.2C163.783 27.2 160.997 28.3917 158.76 30.775C156.56 33.1583 155.46 36.1467 155.46 39.74C155.46 43.3333 156.56 46.3217 158.76 48.705C160.997 51.0883 163.783 52.28 167.12 52.28C170.86 52.28 173.83 50.8133 176.03 47.88H183.73C182.227 51.2533 180.045 53.9117 177.185 55.855C174.362 57.7617 171.007 58.715 167.12 58.715C161.73 58.715 157.238 56.9183 153.645 53.325C150.052 49.695 148.255 45.1667 148.255 39.74C148.255 34.3133 150.052 29.8033 153.645 26.21C157.238 22.58 161.73 20.765 167.12 20.765ZM200.762 31.16C204.759 31.16 207.949 32.095 210.332 33.965C212.752 35.835 213.962 38.4383 213.962 41.775V58H207.032V56.02C205.236 57.8167 202.412 58.715 198.562 58.715C195.336 58.715 192.806 57.945 190.972 56.405C189.176 54.865 188.277 52.6833 188.277 49.86C188.277 46.8533 189.432 44.7267 191.742 43.48C194.089 42.1967 197.041 41.555 200.597 41.555H207.032C207.032 40.125 206.464 39.0617 205.327 38.365C204.191 37.6683 202.724 37.32 200.927 37.32C197.737 37.32 195.702 38.365 194.822 40.455H187.672C188.369 37.4117 189.927 35.1017 192.347 33.525C194.804 31.9483 197.609 31.16 200.762 31.16ZM207.032 47.825V46.725H200.597C197.004 46.725 195.207 47.7883 195.207 49.915C195.207 50.8683 195.611 51.6567 196.417 52.28C197.261 52.8667 198.434 53.16 199.937 53.16C201.917 53.16 203.586 52.7017 204.942 51.785C206.336 50.8317 207.032 49.5117 207.032 47.825ZM221.396 21.205H228.326V58H221.396V21.205ZM248.65 52.335C251.326 52.335 253.27 51.3633 254.48 49.42H261.52C260.566 52.28 258.916 54.5533 256.57 56.24C254.26 57.89 251.62 58.715 248.65 58.715C244.69 58.715 241.371 57.4133 238.695 54.81C236.055 52.17 234.735 48.87 234.735 44.91C234.735 40.95 236.055 37.6683 238.695 35.065C241.371 32.4617 244.69 31.16 248.65 31.16C251.62 31.16 254.26 32.0033 256.57 33.69C258.916 35.34 260.566 37.595 261.52 40.455H254.48C253.27 38.5117 251.326 37.54 248.65 37.54C246.56 37.54 244.873 38.2183 243.59 39.575C242.306 40.9317 241.665 42.71 241.665 44.91C241.665 47.11 242.306 48.9067 243.59 50.3C244.873 51.6567 246.56 52.335 248.65 52.335Z" fill="url(#paint0_linear_5_284)"/>
          <path d="M22.935 21.48L38.83 58H31.185L27.61 49.915H10.945L7.37 58H0L15.895 21.48H22.935ZM19.305 31.05L13.805 43.48H24.75L19.305 31.05ZM68.9859 31.875L57.1609 58H50.6159L38.7909 31.875H46.4359L54.0259 48.705L61.6159 31.875H68.9859ZM73.242 27.64V20.105H80.832V27.64H73.242ZM73.572 31.875H80.502V58H73.572V31.875ZM90.8647 25.99H97.7947V31.875H105.715V37.925H97.7947V48.705C97.7947 50.8683 98.8031 51.95 100.82 51.95H105.715V58H100.765C97.5747 58 95.1181 57.2667 93.3947 55.8C91.7081 54.2967 90.8647 52.2983 90.8647 49.805V37.925H86.0797V31.875H90.8647V25.99ZM113.804 35.12C116.518 32.48 119.909 31.16 123.979 31.16C128.049 31.16 131.423 32.48 134.099 35.12C136.813 37.7233 138.169 41.005 138.169 44.965C138.169 48.8883 136.813 52.17 134.099 54.81C131.423 57.4133 128.049 58.715 123.979 58.715C119.909 58.715 116.518 57.4133 113.804 54.81C111.128 52.17 109.789 48.8883 109.789 44.965C109.789 41.005 111.128 37.7233 113.804 35.12ZM129.149 39.63C127.793 38.2367 126.069 37.54 123.979 37.54C121.889 37.54 120.148 38.2367 118.754 39.63C117.398 41.0233 116.719 42.8017 116.719 44.965C116.719 47.1283 117.398 48.9067 118.754 50.3C120.148 51.6567 121.889 52.335 123.979 52.335C126.069 52.335 127.793 51.6567 129.149 50.3C130.543 48.9067 131.239 47.1283 131.239 44.965C131.239 42.8017 130.543 41.0233 129.149 39.63Z" fill="url(#paint1_linear_5_284)"/>
          <path d="M25.935 21.48L41.83 58H34.185L30.61 49.915H13.945L10.37 58H3L18.895 21.48H25.935ZM22.305 31.05L16.805 43.48H27.75L22.305 31.05ZM70.3359 31.875L58.5109 58H51.9659L40.1409 31.875H47.7859L55.3759 48.705L62.9659 31.875H70.3359ZM72.942 27.64V20.105H80.532V27.64H72.942ZM73.272 31.875H80.202V58H73.272V31.875ZM88.9147 25.99H95.8447V31.875H103.765V37.925H95.8447V48.705C95.8447 50.8683 96.8531 51.95 98.8697 51.95H103.765V58H98.8147C95.6247 58 93.1681 57.2667 91.4447 55.8C89.7581 54.2967 88.9147 52.2983 88.9147 49.805V37.925H84.1297V31.875H88.9147V25.99ZM110.204 35.12C112.918 32.48 116.309 31.16 120.379 31.16C124.449 31.16 127.823 32.48 130.499 35.12C133.213 37.7233 134.569 41.005 134.569 44.965C134.569 48.8883 133.213 52.17 130.499 54.81C127.823 57.4133 124.449 58.715 120.379 58.715C116.309 58.715 112.918 57.4133 110.204 54.81C107.528 52.17 106.189 48.8883 106.189 44.965C106.189 41.005 107.528 37.7233 110.204 35.12ZM125.549 39.63C124.193 38.2367 122.469 37.54 120.379 37.54C118.289 37.54 116.548 38.2367 115.154 39.63C113.798 41.0233 113.119 42.8017 113.119 44.965C113.119 47.1283 113.798 48.9067 115.154 50.3C116.548 51.6567 118.289 52.335 120.379 52.335C122.469 52.335 124.193 51.6567 125.549 50.3C126.943 48.9067 127.639 47.1283 127.639 44.965C127.639 42.8017 126.943 41.0233 125.549 39.63Z" fill="white"/>
          <defs>
          <linearGradient id="paint0_linear_5_284" x1="146" y1="37.5" x2="264" y2="37.5" gradientUnits="userSpaceOnUse">
          <stop stop-color="#00AAFF"/>
          <stop offset="0.305" stop-color="#9054EA"/>
          <stop offset="0.58" stop-color="#FF4053"/>
          <stop offset="1" stop-color="#04E061"/>
          </linearGradient>
          <linearGradient id="paint1_linear_5_284" x1="0" y1="37.5" x2="141" y2="37.5" gradientUnits="userSpaceOnUse">
          <stop stop-color="#00AAFF"/>
          <stop offset="0.305" stop-color="#9054EA"/>
          <stop offset="0.58" stop-color="#FF4053"/>
          <stop offset="1" stop-color="#04E061"/>
          </linearGradient>
          </defs>
        </svg>
        </div>
        <div className="text">
          <p>
          Значимость этих проблем настолько очевидна, что новая модель организационной деятельности требуют от нас анализа форм развития.      </p>
          <br/>
          <p>Таким образом дальнейшее развитие различных форм деятельности способствует подготовки и реализации форм развития.
      
          </p>
        </div>
        <Toolbar />
        <Divider />
        <CollapseReg/>
        {/* <List>
          {['Inbox', 'Starred', 'Send email', 'Drafts'].map((text, index) => (
            <ListItem key={text} disablePadding>
              <ListItemButton>
                <ListItemIcon>
                  {index % 2 === 0 ? <InboxIcon /> : <MailIcon />}
                </ListItemIcon>
                <ListItemText primary={text} />
              </ListItemButton>
            </ListItem>
          ))}
        </List> */}
        <Divider />
        {/* <List>
          {['All mail', 'Trash', 'Spam'].map((text, index) => (
            <ListItem key={text} disablePadding>
              <ListItemButton>
                <ListItemIcon>
                  {index % 2 === 0 ? <InboxIcon /> : <MailIcon />}
                </ListItemIcon>
                <ListItemText primary={text} />
              </ListItemButton>
            </ListItem>
          ))}
        </List> */}
        <Price/>
      </Drawer>
      <div className="adept">
      </div>
      </>
    //   </div>
    //   {/* 
    //   <Field/>
    //   <Box
    //     component="main"
    //     sx={{ flexGrow: 1, bgcolor: 'background.default', p: 3 }}> 
    //   </Box> 
    // </Box> */}
  );
}
