'use strict';

Object.defineProperty(exports, "__esModule", {
  value: true
});

var _react = require('react');

var _react2 = _interopRequireDefault(_react);

var _Input = require('./Input');

var _Input2 = _interopRequireDefault(_Input);

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

var Checkbox = function Checkbox(props) {
  return _react2.default.createElement(
    'div',
    {
      className: 'mdc-checkbox'
    },
    _react2.default.createElement(_Input2.default, props),
    _react2.default.createElement(
      'div',
      { className: 'mdc-checkbox__background' },
      _react2.default.createElement(
        'svg',
        {
          version: '1.1',
          className: 'mdc-checkbox__checkmark',
          xmlns: 'http://www.w3.org/2000/svg',
          viewBox: '0 0 24 24'
        },
        _react2.default.createElement('path', {
          className: 'mdc-checkbox__checkmark__path',
          fill: 'none',
          stroke: 'white',
          d: 'M1.73,12.91 8.1,19.28 22.79,4.59'
        })
      ),
      _react2.default.createElement('div', { className: 'mdc-checkbox__mixedmark' })
    )
  );
};

exports.default = Checkbox;