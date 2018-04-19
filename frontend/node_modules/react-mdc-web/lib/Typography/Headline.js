'use strict';

Object.defineProperty(exports, "__esModule", {
  value: true
});

var _extends = Object.assign || function (target) { for (var i = 1; i < arguments.length; i++) { var source = arguments[i]; for (var key in source) { if (Object.prototype.hasOwnProperty.call(source, key)) { target[key] = source[key]; } } } return target; };

var _react = require('react');

var _react2 = _interopRequireDefault(_react);

var _TypographyElement = require('./TypographyElement');

var _TypographyElement2 = _interopRequireDefault(_TypographyElement);

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

var defaultProps = {
  component: 'h1'
};

var Headline = function Headline(props) {
  return _react2.default.createElement(_TypographyElement2.default, _extends({
    modificator: 'headline'
  }, props));
};

Headline.defaultProps = defaultProps;
exports.default = Headline;